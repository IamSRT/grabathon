package com.grabathon.paymentassistant.storage.db.repository;

import com.grabathon.paymentassistant.storage.db.entity.Template;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface TemplateRepository extends JpaRepository<Template, Long> { }
