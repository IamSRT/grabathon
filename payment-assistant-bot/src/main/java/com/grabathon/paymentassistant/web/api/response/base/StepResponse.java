package com.grabathon.paymentassistant.web.api.response.base;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.grabathon.paymentassistant.web.api.response.vo.StepResponseVO;
import com.grabathon.paymentassistant.web.api.response.vo.TemplateResponseVO;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.util.List;

@Getter @Setter @NoArgsConstructor @AllArgsConstructor
public class StepResponse {

    @JsonProperty ("options") private List<StepResponseVO> steps;
    @JsonProperty ("messages") private List<TemplateResponseVO> templates;

}
