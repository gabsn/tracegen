from ddtrace import tracer
from time import sleep

while True:
    span = tracer.start_span(name='test-python', service='tracegen', resource='python' )
    span.finish()
    sleep(1)

