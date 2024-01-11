package com.newrelic.app;

import java.util.HashMap;
import java.util.Map;
import java.util.Random;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.newrelic.api.agent.NewRelic;

@RestController
@RequestMapping()
public class Controller {

	private static final String CUSTOM_EVENT = "MyCustomEvent";

	private static final Logger logger = LogManager.getLogger(Controller.class);

	private Random r = new Random();

	public Controller() {
	}

	@GetMapping("api")
	public ResponseEntity<String> parentMethod() {

		try {
			logger.info("Main method is triggered...");
			childMethod();
			return new ResponseEntity<String>("Done.", HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<String>("Done.", HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}

	private void childMethod() throws Exception {

		logger.info("Child method method is triggered...");
		if (r.nextInt(15) != 5) {
			logger.info("Method succeeded.");
		} else {
			Map<String, Object> attrs = new HashMap<String, Object>();
			attrs.put("mykey", "myvalue");
			NewRelic.getAgent().getInsights().recordCustomEvent(CUSTOM_EVENT, attrs);

			logger.info("Method failed.");

			throw new Exception("MyException");
		}
	}
}