package com.grabathon.paymentassistant.storage.db.entity;

import lombok.Getter;

import javax.persistence.*;

@Getter
@Entity @Table(name = "step") public class Step {

    @Id @GeneratedValue @Column (name = "id")
    private Long id;

    @Column (name = "title")
    private String title;

    @Column (name = "priority")
    private Integer priority;

    @Column (name = "actions")
    private String actions;

    @Column (name = "templates")
    private String templates;

    @Column (name = "next_steps")
    private String nextSteps;

    @Column (name = "description")
    private String description;

    @Column (name = "render_type")
    private String renderType;

    @Column (name = "rule")
    private String rule;

}
