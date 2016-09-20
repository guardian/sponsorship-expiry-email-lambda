package main

import (
	//"github.com/apex/go-apex"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/aws"
	"html/template"
	"time"
	"fmt"
	"bytes"
)

const emailTemplate = `
<!DOCTYPE html>
<html>
	<body>
		<h3>Expiring sponsorships</h3>
		{{if .Expiring}}
			<table>
				{{range .Expiring}}
					<tr>
					<td>{{with $asset := index .SponsorLogo.Assets 0}}<img src="{{$asset.ImageUrl}}"/>{{end}}</td>
					<td>
						{{.SponsorshipType}}<br />
						Sponsor: <a href="https://tagmanager.gutools.co.uk/sponsorship/{{.Id}}">{{.SponsorName}}</a><br />
						Expires: {{formatMillisDate .ValidTo}}
					</td>
					</tr>
				{{end}}
			</table>
		{{else}}
			<p>no sponsorships expiring in the next 7 days</p>
		{{end}}

		<h3>Expired sponsorships</h3>
		{{if .Expiring}}
			<table>
				{{range .Expiring}}
					<tr>
					<td>{{with $asset := index .SponsorLogo.Assets 0}}<img src="{{$asset.ImageUrl}}"/>{{end}}</td>
					<td>
						{{.SponsorshipType}}<br />
						Sponsor: <a href="https://tagmanager.gutools.co.uk/sponsorship/{{.Id}}">{{.SponsorName}}</a><br />
						Expired: {{formatMillisDate .ValidTo}}
					</td>
					</tr>
				{{end}}
			</table>
		{{else}}
			<p>no sponsorships expired in the last 7 days</p>
		{{end}}
	</body>
</html>`

func formatMillisDate(millis int64) string {
	return time.Unix(millis / 1000, 0).Format(time.RFC1123)
}

func main() {
	awsSession := session.New(&aws.Config{
		Region: aws.String("eu-west-1"),
	})

	dynamo := dynamodb.New(awsSession)
	ses := ses.New(awsSession)

	expiring, err := LoadExpiringSoon(dynamo)

	if (err != nil) {
		return
	}

	expired, err := LoadExpiredRecently(dynamo)

	if (err != nil) {
		return
	}

	t, err := template.New("webpage").Funcs(template.FuncMap{"formatMillisDate": formatMillisDate,}).Parse(emailTemplate)

	templateData := struct {
		Expiring []Sponsorship
		Expired []Sponsorship
	}{
		Expiring: expiring,
		Expired: expired,
	}

	messageBuffer := new(bytes.Buffer)
	err = t.Execute(messageBuffer, templateData)


	if (err != nil) {
		fmt.Println(err.Error())
	}

	err = sendEmail(messageBuffer.String(), ses)

	if (err != nil) {
		fmt.Println(err.Error())
	}

	//apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
	//	dynamo := dynamodb.New(session.New())
	//
	//	LoadExpiringSoon(dynamo)
	//
	//	return m, nil
	//})
}