(Roadmap.sh)
When looking for a website, we type the Domain name: https://wikipedia.com, the browser checks its cache to see 
if:
	ip address == found?
else if:
	the OS checks ram cache to find what ip address it points to.
else:
	we ask the resolver (belongs to ISP)

Resolver checks it's cache, if !found:
	Resolver passes to root (13 root servers in the world, multiple servers in each one)

Root searches the TLD, this query search takes some time... 

Using the TLD domain, you have the registrar, which points to that specific domain (glued addresses).

Resolver passes the ip address to the OS. It's saved in cache. 