package com.grabathon.paymentassistant.web.api.request;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.grabathon.paymentassistant.web.api.request.base.StepsRequest;
import lombok.Getter;
import lombok.Setter;

@Getter @Setter
@JsonIgnoreProperties(ignoreUnknown = true)
public class NextStepsRequest extends StepsRequest {

    public NextStepsRequest () {
        super();
    }

}
