package com.grabathon.paymentassistant.web.api.request.base;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Vouch{
    @JsonProperty("id")
    public String Id;

    @JsonProperty("vouchee_id")
    public String VoucheeId;

    @JsonProperty("voucher_id")
    public String VoucherId;

    @JsonProperty("vouch_type")
    public String VouchType;

    @JsonProperty("amount")
    public Float Amount;
}
