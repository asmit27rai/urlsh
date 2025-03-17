import { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card, CardContent } from "@/components/ui/card";

interface Url {
  short_code: string;
  long_url: string;
  clicks: number;
}

interface ShortenResponse {
  short_code: string;
  long_url: string;
  clicks: number;
}

function App() {
  const [urls, setUrls] = useState<Url[]>([]);
  const [longUrl, setLongUrl] = useState<string>("");
  const Backend_URL = import.meta.env.VITE_BACKEND_URL;

  const fetchUrls = async (): Promise<void> => {
    try {
      const response = await fetch(`${Backend_URL}/urls`);
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      const data: Url[] = await response.json();
      setUrls(data);
    } catch (error) {
      console.error("Error fetching URLs:", error);
    }
  };

  const handleSubmit = async (
    e: React.FormEvent<HTMLFormElement>
  ): Promise<void> => {
    e.preventDefault();
    try {
      const response = await fetch(`${Backend_URL}/shorten`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ long_url: longUrl }),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! Status: ${response.status}`);
      }
      const data: ShortenResponse = await response.json();
      setUrls((prevUrls) => [...prevUrls, data]);
      setLongUrl("");
    } catch (error) {
      console.error("Error shortening URL:", error);
    }
  };

  const trackClick = async (shortCode: string): Promise<void> => {
    try {
      await fetch(`${Backend_URL}/${shortCode}/track`, {
        method: "POST",
      });
      fetchUrls();
    } catch (error) {
      console.error("Error tracking click:", error);
    }
  };

  const handleNavigate = (shortCode: string): void => {
    window.open(`http://localhost:8080/${shortCode}`, "_blank");
  };

  useEffect(() => {
    fetchUrls();
  }, []);

  return (
    <div className="min-h-screen bg-gray-100 py-8 flex justify-center items-center">
      <div className="max-w-2xl w-full px-4 space-y-6">
        <h1 className="text-4xl font-bold text-center">URL Shortener</h1>
        <Card>
          <CardContent className="p-6 space-y-4">
            <form onSubmit={handleSubmit} className="space-y-4">
              <Input
                type="text"
                value={longUrl}
                onChange={(e) => setLongUrl(e.target.value)}
                placeholder="Enter a long URL"
                className="w-full"
              />
              <Button type="submit" className="w-full">
                Shorten
              </Button>
            </form>
          </CardContent>
        </Card>
        <div>
          <h2 className="text-2xl font-semibold mb-4">Shortened URLs</h2>
          {urls.length > 0 ? (
            <div className="space-y-4">
              {urls.map((url, index) => (
                <Card key={index}>
                  <CardContent className="p-4 space-y-2">
                    <div className="flex justify-between items-center">
                      <div>
                        <p className="text-sm text-gray-500">Long URL:</p>
                        <a
                          href={url.long_url}
                          target="_blank"
                          rel="noopener noreferrer"
                          className="text-blue-500 hover:underline break-all"
                        >
                          {url.long_url}
                        </a>
                      </div>
                      <Button
                        size="sm"
                        onClick={() => handleNavigate(url.short_code)}
                      >
                        Navigate
                      </Button>
                    </div>
                    <div>
                      <p className="text-sm text-gray-500">Short URL:</p>
                      <a
                        href={`http://localhost:8080/${url.short_code}`}
                        target="_blank"
                        rel="noopener noreferrer"
                        className="text-blue-500 hover:underline break-all"
                        onClick={() => trackClick(url.short_code)}
                      >
                        {`http://localhost:8080/${url.short_code}`}
                      </a>
                    </div>
                    <p className="text-sm text-gray-500">
                      Clicks: {url.clicks}
                    </p>
                  </CardContent>
                </Card>
              ))}
            </div>
          ) : (
            <p className="text-gray-500 text-center">No shortened URLs yet.</p>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
