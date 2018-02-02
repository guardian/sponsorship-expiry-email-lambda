# Sponsorship Expiry Email Lambda

## Overview
A small lambda which queries the tag manager sponsorship dynamo table to find sponsorships that are expiring, or have expired and sends details of these out in an email.

The labda runs in the composer AWS account and is triggered via the cloudwatch scheduler. The schedule is currently configured to send out the email daily at 7AM GMT.

Updating email recipients
-------------------------

This is the most likely maintenance task. The recipients of the email can be configured in the AWS console within the Lambda: `sponsorship-expiry-email-lambda-STAGE` by modifying the `toAddresses` environment variable. The values passed to this variable should be provided in a comma separated list (without spaces).