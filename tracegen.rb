require 'ddtrace'

loop do
	Datadog.tracer.trace('test-ruby') do |span|
		span.service = 'tracegen'
		span.resource = 'ruby'
	end
	sleep 1
end
