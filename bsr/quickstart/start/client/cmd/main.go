// Copyright 2020-2026 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"log"
	"net/http"

	tagv1 "buf.build/gen/go/aquarium/common/protocolbuffers/go/tag/v1"
	"buf.build/gen/go/aquarium/invoice/connectrpc/gosimple/invoice/v1/invoicev1connect"
	invoicev1 "buf.build/gen/go/aquarium/invoice/protocolbuffers/go/invoice/v1"
)

func main() {
	client := invoicev1connect.NewInvoiceServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)

	_, err := client.CreateInvoice(
		context.Background(),
		&invoicev1.CreateInvoiceRequest{
			Invoice: &invoicev1.Invoice{
				InvoiceId:  "invoice-one",
				CustomerId: "customer-one",
				LineItems: []*invoicev1.LineItem{
					{
						UnitPrice: 999,
						Quantity:  2,
					},
				},
			},
			Tags: &tagv1.Tags{
				Tag: []string{
					"bogo-campaign",
					"valued-customer",
				},
			},
		},
	)
	if err != nil {
		log.Fatalf("error creating valid invoice: %v", err)
	}
	log.Println("Valid invoice created")
}
