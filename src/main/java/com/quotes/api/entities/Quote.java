package com.quotes.api.entities;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.format.annotation.DateTimeFormat;

import java.sql.Timestamp;
import java.util.Date;

@Document(collection = "quotes")
public class Quote {
    @Id
    private String id;
    private String author;
    private String work;
    private String message;
    private String[] tags;

    @DateTimeFormat(iso = DateTimeFormat.ISO.DATE_TIME)
    private Date dateCreated;

    public String getAuthor() {
        return author;
    }

    public String getId() { return id; }

    public void setAuthor(String author) {
        this.author = author;
    }

    public String getWork() {
        return work;
    }

    public void setWork(String work) {
        this.work = work;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public String[] getTags() {
        return tags;
    }

    public void setTags(String[] tags) {
        this.tags = tags;
    }

    public Date getDateCreated() {
        return dateCreated;
    }

    public void setDateCreated() {
        this.dateCreated = new Timestamp(new Date().getTime());
    }
}
