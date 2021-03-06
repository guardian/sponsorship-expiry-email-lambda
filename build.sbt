organization  := "com.gu"
description   := "AWS Lambda to send email to stakeholders for expired sponsorships."
scalaVersion  := "2.12.4"
scalacOptions ++= Seq("-deprecation", "-feature", "-unchecked", "-target:jvm-1.8", "-Xfatal-warnings")
name := "sponsorship-expiry-email-lambda"

enablePlugins(SbtTwirl, JavaAppPackaging, RiffRaffArtifact)

val AwsSdkVersion = "1.11.269"

libraryDependencies ++= Seq(
  "com.amazonaws" % "aws-java-sdk-lambda" % AwsSdkVersion,
  "com.amazonaws" % "aws-java-sdk-sts" % AwsSdkVersion,
  "com.amazonaws" % "aws-java-sdk-ses" % AwsSdkVersion,
  "com.squareup.okhttp3" % "okhttp" % "3.6.0",
  "com.gu" %% "scanamo" % "1.0.0-M4"
)

 topLevelDirectory in Universal := None
 packageName in Universal := normalizedName.value

 riffRaffManifestProjectName := s"Editorial Tools::${name.value}"
 riffRaffPackageName := "sponsorship-expiry-email-lambda"
 riffRaffPackageType := (packageBin in Universal).value
 riffRaffUploadArtifactBucket := Option("riffraff-artifact")
 riffRaffUploadManifestBucket := Option("riffraff-builds")

TwirlKeys.templateImports += "com.gu.comdev.sponsorshipexpiry.models._"
TwirlKeys.templateImports += "org.joda.time.format.DateTimeFormat"
TwirlKeys.templateImports += "org.joda.time.DateTime"

initialize := {
  val _ = initialize.value
  assert(sys.props("java.specification.version") == "1.8",
    "Java 8 is required for this project.")
}