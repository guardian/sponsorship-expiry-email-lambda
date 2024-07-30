organization  := "com.gu"
description   := "AWS Lambda to send email to stakeholders for expired sponsorships."
scalaVersion  := "2.12.19"
scalacOptions ++= Seq("-deprecation", "-feature", "-unchecked", "-target:jvm-1.8", "-Xfatal-warnings")
name := "sponsorship-expiry-email-lambda"

enablePlugins(SbtTwirl, JavaAppPackaging)

val AwsSdkVersion = "1.12.666"

libraryDependencies ++= Seq(
  "com.amazonaws" % "aws-java-sdk-lambda" % AwsSdkVersion,
  "com.amazonaws" % "aws-java-sdk-sts" % AwsSdkVersion,
  "com.amazonaws" % "aws-java-sdk-ses" % AwsSdkVersion,
  "com.amazonaws" % "aws-java-sdk-s3" % AwsSdkVersion,
  "com.squareup.okhttp3" % "okhttp" % "4.9.3",
  "com.gu" %% "scanamo" % "1.0.0-M8"
)

Universal / topLevelDirectory := None
Universal / packageName := normalizedName.value

dependencyOverrides += "org.jetbrains.kotlin" % "kotlin-stdlib" % "[1.6.0,)"

TwirlKeys.templateImports += "com.gu.comdev.sponsorshipexpiry.models._"
TwirlKeys.templateImports += "org.joda.time.format.DateTimeFormat"
TwirlKeys.templateImports += "org.joda.time.DateTime"

initialize := {
  val _ = initialize.value
  assert(sys.props("java.specification.version").toInt >= 11,
    "Java 11 is required for this project.")
}
