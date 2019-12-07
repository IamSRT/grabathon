package com.grabathon.paymentassistant.web.api.request.base;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.util.List;

@JsonInclude(JsonInclude.Include.NON_NULL)
@Getter @Setter @NoArgsConstructor @AllArgsConstructor
public class StepsRequest {
    @JsonProperty("user_id")
    public String UserId;

    @JsonProperty("vouches")
    public List<Vouch> Vouches;

    @JsonProperty("vouch")
    public Vouch Vouch;

    @JsonProperty("payment_id")
    public String PaymentId;

    @JsonProperty("payment")
    public Payment payment;

    public boolean isRequester;
    public String requesterNumber;
    public List<String> requesteeNumber;
    public Float amount;
    public Integer vouch_id;


}

