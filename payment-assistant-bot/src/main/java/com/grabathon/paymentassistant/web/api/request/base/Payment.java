package com.grabathon.paymentassistant.web.api.request.base;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * Created by Sai Ravi Teja K on 08, Dec 2019
 * © Bundl Technologies Private Ltd.
 */

public class Payment {
    @JsonProperty("id")
    public String Id;

    @JsonProperty("status")
    public String Status;

    @JsonProperty("type")
    public String Type;

    @JsonProperty("transactions")
    public Transaction[] Transactions;
}
