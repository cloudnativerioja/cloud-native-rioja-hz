package main

import (
	"fmt"
	"os"

	"github.com/pulumi/pulumi-cloudflare/sdk/v5/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	// Get the zone ID from the environment variables
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	if zoneID != "" {
		fmt.Println("ZoneID is set")
	} else {
		fmt.Println("ZoneID is NOT set")
		os.Exit(2)
	}

	// Start the Pulumi program
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Replace `example.com` with your domain and `zoneId` with your zone ID.
		zoneId := pulumi.String(zoneID)
		domain := pulumi.String("cloudnativerioja.com")

		// Create the first CNAME record
		_, err := cloudflare.NewRecord(ctx, "root", &cloudflare.RecordArgs{
			Name:    pulumi.Sprintf("%s", domain),
			Type:    pulumi.String("CNAME"),
			Value:   pulumi.String("cloudnativerioja.github.io"),
			ZoneId:  zoneId,
			Proxied: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		// Create the second CNAME record
		_, err = cloudflare.NewRecord(ctx, "docs", &cloudflare.RecordArgs{
			Name:    pulumi.Sprintf("docs.%s", domain),
			Type:    pulumi.String("CNAME"),
			Value:   pulumi.String("cloudnativerioja.github.io"),
			ZoneId:  zoneId,
			Proxied: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		// Create the third CNAME record
		_, err = cloudflare.NewRecord(ctx, "www", &cloudflare.RecordArgs{
			Name:    pulumi.Sprintf("www.%s", domain),
			Type:    pulumi.String("CNAME"),
			Value:   pulumi.String("cloudnativerioja.github.io"),
			ZoneId:  zoneId,
			Proxied: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		// Create the A record
		_, err = cloudflare.NewRecord(ctx, "local-cfp", &cloudflare.RecordArgs{
			Name:    pulumi.Sprintf("cfp.%s", domain),
			Type:    pulumi.String("A"),
			Value:   pulumi.String("192.0.2.1"),
			ZoneId:  zoneId,
			Proxied: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		// Create the A record
		_, err = cloudflare.NewRecord(ctx, "permission-manager", &cloudflare.RecordArgs{
			Name:    pulumi.Sprintf("perms.%s", domain),
			Type:    pulumi.String("A"),
			Value:   pulumi.String("74.220.16.39"),
			ZoneId:  zoneId,
			Proxied: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		// Create the TXT record
		_, err = cloudflare.NewRecord(ctx, "txt-github", &cloudflare.RecordArgs{
			Name:   pulumi.Sprintf("_github-challenge-cloudnativerioja-org"),
			Type:   pulumi.String("TXT"),
			Value:  pulumi.String("f7dbc6f841"),
			ZoneId: zoneId,
		})
		if err != nil {
			return err
		}

		_, err = cloudflare.NewRecord(ctx, "txt-github-pages", &cloudflare.RecordArgs{
			Name:   pulumi.Sprintf("_github-pages-challenge-cloudnativerioja"),
			Type:   pulumi.String("TXT"),
			Value:  pulumi.String("f8e61a9a46363cd21f945152afca41"),
			ZoneId: zoneId,
		})
		if err != nil {
			return err
		}

		// Create the A record
		_, err = cloudflare.NewRecord(ctx, "ns1-civo", &cloudflare.RecordArgs{
			Name:    pulumi.Sprintf("cluster.%s", domain),
			Type:    pulumi.String("NS"),
			Value:   pulumi.String("ns0.civo.com"),
			ZoneId:  zoneId,
		})
		if err != nil {
			return err
		}

		// Create the A record
		_, err = cloudflare.NewRecord(ctx, "ns2-civo", &cloudflare.RecordArgs{
			Name:    pulumi.Sprintf("cluster.%s", domain),
			Type:    pulumi.String("NS"),
			Value:   pulumi.String("ns1.civo.com"),
			ZoneId:  zoneId,
		})
		if err != nil {
			return err
		}

		return nil
	})
}
