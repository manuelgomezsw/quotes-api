package com.quotes.api.controllers;

import com.quotes.api.entities.Quote;
import com.quotes.api.repository.QuoteRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api")
public class QuoteController {
    @Autowired
    QuoteRepository quoteRepository;

    @GetMapping("/quotes")
    public ResponseEntity<List<Quote>> getAllQuotes() {
        try {
            List<Quote> quotes = quoteRepository.findAll();
            return new ResponseEntity<>(quotes, HttpStatus.OK);
        }
        catch (Exception exception) {
            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @PostMapping("/quotes")
    public ResponseEntity<Quote> create(@RequestBody Quote quote) {
        try {
            quote.setDateCreated();
            Quote quoteCreated = quoteRepository.save(quote);

            return new ResponseEntity<>(quoteCreated, HttpStatus.CREATED);
        }
        catch (Exception exception) {
            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }
}
