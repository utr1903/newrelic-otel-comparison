package com.otel.app;

import java.util.Random;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.common.AttributeKey;
import io.opentelemetry.api.common.Attributes;
import io.opentelemetry.api.metrics.LongCounter;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;

@RestController
@RequestMapping()
public class Controller {

	private static final String CUSTOM_EVENT = "MyCustomEvent";

	private static final Logger logger = LogManager.getLogger(Controller.class);

	private Tracer tracer;

	private final LongCounter invocations;

	private Random r = new Random();

	public Controller() {

		this.tracer = GlobalOpenTelemetry.getTracer(Controller.class.getName());
		this.invocations = GlobalOpenTelemetry.getMeter(Controller.class.getName())
				.counterBuilder("invocations")
				.setDescription("Measures the number of method invocations.")
				.build();
	}

	@GetMapping("api")
	public ResponseEntity<String> parentMethod() {

		try {
			logger.info("Main method is triggered...");
			childMethod();
			return new ResponseEntity<String>("Done.", HttpStatus.OK);
		} catch (Exception e) {
			Span.current().recordException(e);
			return new ResponseEntity<String>("Done.", HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}

	private void childMethod() throws Exception {
		Span childSpan = tracer.spanBuilder(Controller.class.getName() + ".childMethod").startSpan();
		try (Scope scope = childSpan.makeCurrent()) {

			logger.info("Child method method is triggered...");
			if (r.nextInt(15) != 5) {
				logger.info("Method succeeded.");
				invocations.add(1, Attributes.of(AttributeKey.booleanKey("succeded"), true));
			} else {
				logger.info("Method failed.");
				invocations.add(1, Attributes.of(AttributeKey.booleanKey("succeded"), false));

				Attributes attrs = Attributes.of(AttributeKey.stringKey("mykey"), "myvalue");
				childSpan.addEvent(CUSTOM_EVENT, attrs);

				throw new Exception("MyException");
			}
		} finally {
			childSpan.end();
		}
	}
}