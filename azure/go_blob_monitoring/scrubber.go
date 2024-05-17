package goblobmonitoring

class dcd {
    constructor(name, pattern, replacement) {
        this.name = name;
        this.replacement = replacement;
        this.regexp = RegExp(pattern, 'g');
    }
}

type scrubberRule struct{
	Name 		string
    Replacement string
    Regexp      *regexp.Regexp
}

// Scrubber interface wraps around the Event Platform client
type Scrubber interface {
	newScrubberRule(context context.Context, configs []SCRUBBER_RULE_CONFIGS)
	scrub(batchData interface{}) (err error)
}


func (Scrubber) scrub(record) {
	if (!this.rules) {
		return record;
	}
	this.rules.forEach(rule => {
		record = record.replace(rule.regexp, rule.replacement);
	});
	return record;
}

func (Scrubber) newScrubberRule(context context.Context, configs []SCRUBBER_RULE_CONFIGS) {
	var rules = []
	for name, config := range configs{
			rule:= &scrubberRule{
				Name: name,
				Replacement: config.Replacement,
				Regexp: regexp.MustCompile(pattern),
			}	
			rules.append
			context.log.error(
				`Regexp for rule ${name} pattern ${settings['pattern']} is malformed, skipping. Please update the pattern for this rule to be applied.`
			)
		}

	this.rules = rules;
}