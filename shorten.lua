wrk.method = "GET"
wrk.path = "/expand?shorten={shortenid}"

function request()
	local requestPath = string.gsub(wrk.path,"{shortenid}",math.randm(1,10000))
	return wrk.format(nil,requestPath)
end
