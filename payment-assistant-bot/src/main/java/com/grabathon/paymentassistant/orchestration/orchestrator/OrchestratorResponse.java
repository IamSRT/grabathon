package com.grabathon.paymentassistant.orchestration.orchestrator;

import com.grabathon.paymentassistant.web.api.response.vo.base.StepResponseVO;
import com.grabathon.paymentassistant.web.api.response.vo.base.TemplateResponseVO;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;

import java.util.List;

@Getter @NoArgsConstructor @AllArgsConstructor
public class OrchestratorResponse {

    private List<StepResponseVO> steps;
    private List<TemplateResponseVO> templates;

}
