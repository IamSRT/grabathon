package com.grabathon.paymentassistant.storage.db.wrapper;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;

import java.util.List;

@Getter @NoArgsConstructor @AllArgsConstructor
public class Step {

    private Long id;

    private String title;

    private Integer priority;

    private List<Action> actions;

    private List<Template> templates;

    private String nextSteps;

    private String description;

    private String renderType;

    private String rule;

}
