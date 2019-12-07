package com.grabathon.paymentassistant.web.api.response.vo;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.grabathon.paymentassistant.storage.db.wrapper.Template;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;

@Getter @NoArgsConstructor @AllArgsConstructor
public class TemplateResponseVO {

    @JsonProperty ("message") private String template;
    @JsonProperty ("messageType") private String templateType;


    public TemplateResponseVO (Template template) {
        this.template = template.getTemplate();
        this.templateType = template.getTemplateType();
    }

}
