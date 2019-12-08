package com.grabathon.paymentassistant.web.api.response.base;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.grabathon.paymentassistant.web.api.response.Requestee;
import com.grabathon.paymentassistant.web.api.response.Requester;
import lombok.*;

@Getter @Setter @NoArgsConstructor @AllArgsConstructor @ToString
public class StepResponse {

    private Requester requester;
    @JsonProperty ("requestees") private Requestee requestee;
    private Object data;

}
