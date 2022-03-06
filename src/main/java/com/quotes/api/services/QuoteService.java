package com.quotes.api.services;

import com.quotes.api.entities.Quote;
import com.quotes.api.exceptions.QuoteNotFoundException;
import com.quotes.api.repositories.QuoteRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class QuoteService {
    @Autowired QuoteRepository quoteRepository;

    public List<Quote> getAll(Optional<Integer> page, Optional<Integer> size, String sortOrder) {
    Pageable criteria =
        PageRequest.of(
            page == null ? 0 : page.get(),
            size == null ? 20 : size.get(),
            Sort.by(getSortDirection(sortOrder), "dateCreated"));
        return quoteRepository.findAll(criteria).getContent();
    }

    public List<Quote> findByTag(String tag) {
        return quoteRepository.findByTagsContaining(tag);
    }

    public Quote findById(String idQuote) {
        Optional<Quote> quote = quoteRepository.findById(idQuote);

        if (quote.isEmpty()) {
            throw new QuoteNotFoundException(String.format("Quote with id %s not found", idQuote));
        }

        return quote.get();
    }

    public Quote create(Quote quote) {
        quote.setDateCreated();
        for (short indexTag = 0; indexTag < quote.getTags().length; indexTag++) {
            quote.getTags()[indexTag].toLowerCase();
        }

        return quoteRepository.save(quote);
    }

    public void delete(String idQuote) {
        if (quoteRepository.existsById(idQuote)) {
            quoteRepository.deleteById(idQuote);
        } else {
            throw new QuoteNotFoundException("Quote not found");
        }
    }

    public Quote update(Quote quote) {
        if (quote.getId() == null || quote.getId().trim().length() == 0) {
            throw new QuoteNotFoundException("Quote not found");
        }

        if (quoteRepository.existsById(quote.getId())) {
            quote.setDateCreated();
            return quoteRepository.save(quote);
        } else {
            throw new QuoteNotFoundException("Quote not found");
        }
    }

    private Sort.Direction getSortDirection(String sortOrder) {
        switch (sortOrder) {
            case "ASC":
                return Sort.Direction.ASC;
            default:
                return Sort.Direction.DESC;
        }
    }
}
