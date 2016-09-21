Sponsorship Expiry Email Lambda
===============================

A small lambda which queries the tag manager sponsorship dynamo table to find sponsorships that are expiring, or have expired soon
and send details of these out in an email.

The labda runs in the composer AWS account and is triggered via the cloudwatch scheduler. The schedule is currently configured to send 
out the email daily at 7AM GMT.

Build and deploy
----------------

This lambda is build in go and uses [apex](apex.run) to bundle the binary and node shim. To create a deployable .zip run:

`apex build sendExpiryEmail > sendExpiryEmail.zip`

And then upload the resulting .zip file to lambda via the console. Due to limitations with the federated identity credentials produced
by janus full management via apex didn't work.

Testing
-------

run `go test` in the sendExpiryEmail sub directory. Tests are limited to querying dynamo as we don't want to actually send emails
from the test suite. You will need the details of the composer aws account available to the default provider chain (you can use the
export to shell feature in janus)

Updating email recipients
-------------------------

This is the most likely maintenance task. The recipients of the email are configured in the emailSender.go file, go may not be
familiar to everyone but hopefully updating the to and cc lists should be obvious to anyone.