package com.grabathon.paymentassistant.web.api.request.base;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * Created by Sai Ravi Teja K on 08, Dec 2019
 * © Bundl Technologies Private Ltd.
 */

public class Transaction {
    @JsonProperty("id")
    public String Id;

    @JsonProperty("source_id")
    public String SourceId;

    @JsonProperty("destination_id")
    public String DestinationId;

    @JsonProperty("transaction_type")
    public String TransactionType;

    @JsonProperty("status")
    public String Status;

    @JsonProperty("amount")
    public Float Amount;
}
