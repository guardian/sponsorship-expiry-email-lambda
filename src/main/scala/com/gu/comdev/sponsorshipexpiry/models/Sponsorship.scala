package com.gu.comdev.sponsorshipexpiry.models

case class Sponsorship(
  id: Int,
  validFrom: Long,
  validTo: Long,
  status: String,
  sponsorshipType: String,
  sponsorName: String,
  sponsorLogo: Image,
  tags: Seq[Int])
