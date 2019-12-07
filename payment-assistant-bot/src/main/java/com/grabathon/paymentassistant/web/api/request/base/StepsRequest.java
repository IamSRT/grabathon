package com.grabathon.paymentassistant.web.api.request.base;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

@JsonInclude(JsonInclude.Include.NON_NULL)
public class StepsRequest {
    @JsonProperty("user_id")
    public String UserId;

    @JsonProperty("vouches")
    public Vouch[] Vouches;

    @JsonProperty("vouch")
    public Vouch[] Vouch;

    @JsonProperty("payment_id")
    public String PaymentId;

    @JsonProperty("payment")
    public Payment payment;
}

