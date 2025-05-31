export const mapRequest = async (requestJson, requestMapper) => {
	const res = await fetch("http://localhost:8080/api/map-request", {
		method: "POST",
		headers: { "Content-Type": "application/json" },
		body: JSON.stringify({
			requestJson: JSON.parse(requestJson),
			requestMapper: JSON.parse(requestMapper),
		}),
	});
	return res;
};
