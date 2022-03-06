package com.quotes.api.repositories;

import com.quotes.api.entities.Quote;
import org.springframework.data.mongodb.repository.MongoRepository;

import java.util.List;

public interface QuoteRepository extends MongoRepository<Quote, String> {
    List<Quote> findByTagsContaining(String tag);
}