ok, for the time being, lets just implement http 1.0 for now,
i wanted to start from 0.9 but unfortunately, most browsers dont
support it now and we might not be able to test it properly.

so we will go with http 1.0 for now

specs -
pretty simple actually, since http 1.0 introduced support for media types also,
1.
first we need to specify the protocol we are using to send the data
2.
we need to send the content header to tell the type of the file we are sending
like text, html, css or images or something.
3.
then we also need to tell the size of the stuff we are sending in bytes

after these things, we can just directly send the stuff that we wanna send in what
ever format we specified in the content header type, maybe 8 bit encoded strings (utf-8)
or ascii or some weird ass format, doesn't matter tbh, depends on the end consumer.

since http1.0 works on tcp, we will also listen on tcp (i am not implementing tcp here)

# Issue 1 -
how do i know when to stop waiting for data and instead respond back?
well easy, for GET requests sent by browser, the header always ends with \r\n\r\n
we need to use that