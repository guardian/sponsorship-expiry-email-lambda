package com.gu.comdev.sponsorshipexpiry

import com.amazonaws.auth.DefaultAWSCredentialsProviderChain
import com.amazonaws.regions.Regions
import com.amazonaws.services.dynamodbv2.{ AmazonDynamoDB, AmazonDynamoDBClientBuilder }
import com.gu.comdev.sponsorshipexpiry.models.Sponsorship
import com.gu.scanamo.error.DynamoReadError
import com.gu.scanamo.query.{ AndCondition, Between, Condition }
import com.gu.scanamo.{ Scanamo, Table }
import org.joda.time.DateTime
import com.gu.scanamo.syntax._

object SponsorshipRepository {

  type FilterCondition = AndCondition[(Symbol, String), Condition[Between[Long]]]

  private val dynamoClient: AmazonDynamoDB = AmazonDynamoDBClientBuilder.standard().withRegion(Regions.EU_WEST_1).withCredentials(new DefaultAWSCredentialsProviderChain).build()

  private val sponsorshipsTable = Table[Sponsorship]("tag-manager-sponsorships-PROD")

  def loadExpiringSoon(): List[Sponsorship] = {
    val now = DateTime.now()
    val plusOneWeek = now.plusWeeks(1)

    val withFilter: FilterCondition = Condition('status -> "active") and Condition('validTo between (now.getMillis and plusOneWeek.getMillis))
    executeQuery(withFilter)
  }

  def loadExpiredRecently(): List[Sponsorship] = {
    val now = DateTime.now()
    val minusOneWeek = now.minusWeeks(1)

    val withFilter: FilterCondition = Condition('status -> "expired") and Condition('validTo between (minusOneWeek.getMillis and now.getMillis))
    executeQuery(withFilter)
  }

  private def executeQuery(filterCondition: FilterCondition): List[Sponsorship] = {
    val results: List[Either[DynamoReadError, Sponsorship]] = Scanamo.exec(dynamoClient)(sponsorshipsTable.filter(filterCondition).scan)

    results.flatMap {
      case Left(error) =>
        println(error); None
      case Right(sponsorship) => Some(sponsorship)
    }.sortWith(_.validTo < _.validTo)

  }

}
