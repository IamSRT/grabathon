package com.grabathon.paymentassistant.storage.db.wrapper;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter @Setter
@NoArgsConstructor @AllArgsConstructor
public class Template {

    private Long id;

    private String name;

    private String template;

    private String templateType;



    public Template (com.grabathon.paymentassistant.storage.db.entity.Template template) {
        this.id = template.getId();
        this.name = template.getName();
        this.template = template.getTemplate();
        this.templateType = template.getTemplateType();
    }

    public Template (Template template) {
        this.id = template.getId();
        this.name = template.getName();
        this.template = template.getTemplate();
        this.templateType = template.getTemplateType();
    }

}
