package com.grabathon.paymentassistant.storage.db.entity;

import lombok.Getter;

import javax.persistence.*;

@Getter
@Entity @Table(name = "template") public class Template {

    @Id @GeneratedValue @Column (name = "id")
    private Long id;

    @Column (name = "name")
    private String name;

    @Column (name = "template")
    private String template;

    @Column (name = "template_type")
    private String templateType;

}
