import { useState } from "react";
import CodeMirror from "@uiw/react-codemirror";
import { json } from "@codemirror/lang-json";
import { mapRequest } from "./services/api";

const App = () => {
	const [requestJsonStr, setRequestJsonStr] = useState(""); 
	const [requestMapperStr, setRequestMapperStr] = useState(""); 
	const [mappedResult, setMappedResult] = useState("");
	const [loading, setLoading] = useState(false);
	const [error, setError] = useState("");

	// track the last submitted versions for change detection
	const [lastSubmittedRequestJson, setLastSubmittedRequestJson] = useState("");
	const [lastSubmittedRequestMapper, setLastSubmittedRequestMapper] = useState(
		""
	);

	const handleMapRequest = async () => {
		// Validation: Ensure both inputs are non-empty
		if (!requestJsonStr.trim() || !requestMapperStr.trim()) {
			setError("Please provide both Request JSON and Request Mapper.");
			return;
		}

		// here first check if input has changed since last submit
		if (
			requestJsonStr.trim() === lastSubmittedRequestJson.trim() &&
			requestMapperStr.trim() === lastSubmittedRequestMapper.trim()
		) {
			setError("No changes detected since last mapping.");
			return;
		}

		try {
			setLoading(true);
			setError("");
			setMappedResult("");

			const response = await mapRequest(requestJsonStr, requestMapperStr);

			if (!response.ok) {
				throw new Error(`API Error: ${response.statusText}`);
			}

			const data = await response.json();
			setMappedResult(JSON.stringify(data, null, 2));

			// update last submitted values
			setLastSubmittedRequestJson(requestJsonStr.trim());
			setLastSubmittedRequestMapper(requestMapperStr.trim());
		} catch (err) {
			setError(err.message || "An error occurred");
		} finally {
			setLoading(false);
		}
	};

	return (
		<div className="p-4 flex flex-col gap-4">
			<h1 className="text-2xl font-bold">Dynamic JSON Mapper</h1>

			<div className="grid grid-cols-1 md:grid-cols-2 gap-4">
				<div>
					<h2 className="font-semibold">Request JSON</h2>
					<CodeMirror
						value={requestJsonStr}
						height="300px"
						extensions={[json()]}
						onChange={(val) => setRequestJsonStr(val)}
						theme="light"
					/>
				</div>

				<div>
					<h2 className="font-semibold">Request Mapper</h2>
					<CodeMirror
						value={requestMapperStr}
						height="300px"
						extensions={[json()]}
						onChange={(val) => setRequestMapperStr(val)}
						theme="light"
					/>
				</div>
			</div>

			<button
				className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 w-48"
				onClick={handleMapRequest}
				disabled={loading}
			>
				{loading ? "Mapping..." : "Map JSON"}
			</button>

			{error && <div className="text-red-600">Error: {error}</div>}

			{mappedResult && (
				<div>
					<h2 className="font-semibold mt-4">Mapped Output</h2>
					<CodeMirror
						value={mappedResult}
						height="300px"
						extensions={[json()]}
						theme="light"
						readOnly
					/>
				</div>
			)}
		</div>
	);
};

export default App;
