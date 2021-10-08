package com.quotes.api.controllers;

import com.quotes.api.entities.Quote;
import com.quotes.api.repository.QuoteRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;

@RestController
@CrossOrigin(origins = "http://localhost:4200")
@RequestMapping("/api")
public class QuoteController {
    @Autowired
    QuoteRepository quoteRepository;

    @GetMapping("/quotes")
    public ResponseEntity<List<Quote>> GetAll() {
        try {
            List<Quote> quotes = quoteRepository.findAll();
            return new ResponseEntity<>(quotes, HttpStatus.OK);
        }
        catch (Exception exception) {
            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @GetMapping("/quotes/{idQuote}")
    public ResponseEntity<Optional<Quote>> FindById(@PathVariable("idQuote") String idQuote) {
        try {
            Optional<Quote> quote = quoteRepository.findById(idQuote);
            return new ResponseEntity<>(quote, HttpStatus.OK);
        }
        catch (Exception exception) {
            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @PostMapping("/quotes")
    public ResponseEntity<Quote> Create(@RequestBody Quote quote) {
        try {
            quote.setDateCreated();
            Quote quoteCreated = quoteRepository.save(quote);

            return new ResponseEntity<>(quoteCreated, HttpStatus.CREATED);
        }
        catch (Exception exception) {
            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @DeleteMapping("/quotes/{idQuote}")
    public ResponseEntity<String> Delete(@PathVariable("idQuote") String idQuote) {
        try {
            if (quoteRepository.existsById(idQuote)) {
                quoteRepository.deleteById(idQuote);
            } else {
                return new ResponseEntity<>("", HttpStatus.NOT_FOUND);
            }

            return new ResponseEntity<>("Quote Deleted", HttpStatus.OK);
        }
        catch (Exception exception) {
            return new ResponseEntity<>(exception.getMessage(), HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @PutMapping("/quotes")
    public ResponseEntity<Quote> Update(@RequestBody Quote quoteToUpdate) {
        try {
            if (quoteToUpdate.getId() == null || quoteToUpdate.getId().trim().length() == 0) {
                return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
            }

            Quote quoteUpdated;
            if (quoteRepository.existsById(quoteToUpdate.getId())) {
                quoteToUpdate.setDateCreated();
                quoteUpdated = quoteRepository.save(quoteToUpdate);
            } else {
                return new ResponseEntity<>(null, HttpStatus.NOT_FOUND);
            }

            return new ResponseEntity<>(quoteUpdated, HttpStatus.OK);
        }
        catch (Exception exception) {
            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }
}
