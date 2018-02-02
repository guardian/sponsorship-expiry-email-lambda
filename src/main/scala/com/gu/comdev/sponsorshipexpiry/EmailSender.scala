package com.gu.comdev.sponsorshipexpiry

import com.amazonaws.services.simpleemail.model._
import com.amazonaws.auth.DefaultAWSCredentialsProviderChain
import com.amazonaws.regions.Regions
import com.amazonaws.services.simpleemail.{ AmazonSimpleEmailService, AmazonSimpleEmailServiceClientBuilder }
import org.joda.time.DateTime
import org.joda.time.format.DateTimeFormat

import scala.collection.JavaConverters._
import scala.util.{ Failure, Success, Try }

object EmailSender {

  private val sesClient: AmazonSimpleEmailService = AmazonSimpleEmailServiceClientBuilder.standard()
    .withCredentials(new DefaultAWSCredentialsProviderChain)
    .withRegion(Regions.EU_WEST_1).build()

  private val DateFormat = DateTimeFormat.forPattern("E, d MMM y")

  def sendSponsorshipExpiryEmail(toAddresses: Seq[String], ccAddresses: Seq[String], fromAddress: String): Unit = {

    val expired = SponsorshipRepository.loadExpiredRecently()
    val expiring = SponsorshipRepository.loadExpiringSoon()
    val date = DateTime.now().toString(DateFormat)

    val emailBody = html.sponsorshipExpiryEmailBody.apply(expiring, expired).body

    val request = new SendEmailRequest()
      .withDestination(
        new Destination(toAddresses.asJava).withCcAddresses(ccAddresses.asJava)).withMessage(
          new Message()
            .withSubject(new Content().withData("Expiring Advertisement Features " + date))
            .withBody(new Body().withHtml(new Content().withData(emailBody)))).withSource(fromAddress)

    Try(sesClient.sendEmail(request)) match {
      case Success(_) => println(s"Successfully sent sponsorship expiry email. There are ${expired.size} expired sponsorships" +
        s" and ${expiring.size} sponsorships about to expire.")
      case Failure(e) => println(s"Could not send sponsorship expiry email - something went wrong: ${e.getMessage}")
    }
  }

}
