package com.gu.comdev.sponsorshipexpiry

object Lambda {

  private val toAddresses = sys.env.get("toAddresses").map(_.split(",")) getOrElse sys.error("Lambda will not run. Did not receive destination email addresses.")
  private val ccAddresses = sys.env.get("ccAddresses").map(_.split(",")) getOrElse sys.error("Lambda will not run. Did not receive cc addresses.")
  private val fromAddress = sys.env.getOrElse("fromAddress", sys.error("Lambda will not run. Did not receive source email address."))

  def handleRequest(): Unit = {
    EmailSender.sendSponsorshipExpiryEmail(toAddresses, ccAddresses, fromAddress)
  }
}

