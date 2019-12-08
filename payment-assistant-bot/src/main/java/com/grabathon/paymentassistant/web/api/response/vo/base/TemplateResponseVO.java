package com.grabathon.paymentassistant.web.api.response.vo.base;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.grabathon.paymentassistant.storage.db.wrapper.Template;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.ToString;

@Getter @NoArgsConstructor @AllArgsConstructor @ToString
public class TemplateResponseVO {

    @JsonProperty ("message") private String template;
    @JsonProperty ("messageType") private String templateType;


    public TemplateResponseVO (Template template) {
        this.template = template.getTemplate();
        this.templateType = template.getTemplateType();
    }

}
