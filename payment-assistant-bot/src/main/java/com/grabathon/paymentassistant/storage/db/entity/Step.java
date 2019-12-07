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

    @Column (name = "requester_actions")
    private String requesterActions;

    @Column (name = "requestee_actions")
    private String requesteeActions;

    @Column (name = "requester_templates")
    private String requesterTemplates;

    @Column (name = "requestee_templates")
    private String requesteeTemplates;

    @Column (name = "next_steps")
    private String nextSteps;

    @Column (name = "description")
    private String description;

    @Column (name = "requester_render_type")
    private String requesterRenderType;

    @Column (name = "requestee_render_type")
    private String requesteeRenderType;

    @Column (name = "rule")
    private String rule;

}
