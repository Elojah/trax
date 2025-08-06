import { ref } from 'vue';

export interface GeocodingResult {
  place_id: number;
  licence: string;
  osm_type: string;
  osm_id: number;
  lat: string;
  lon: string;
  class: string;
  type: string;
  place_rank: number;
  importance: number;
  addresstype: string;
  name: string;
  display_name: string;
  boundingbox: [string, string, string, string];
}

export const useGeocoding = () => {
  const loading = ref(false);
  const error = ref<string | null>(null);
  const results = ref<GeocodingResult[]>([]);

  // Rate limiting: Nominatim has a usage policy of max 1 request per second
  let lastRequestTime = 0;
  const MIN_INTERVAL = 1000; // 1 second

  const geocode = async (query: string, limit: number = 5): Promise<GeocodingResult[]> => {
    if (!query.trim()) {
      results.value = [];
      return [];
    }

    // Respect rate limiting
    const now = Date.now();
    const timeSinceLastRequest = now - lastRequestTime;
    if (timeSinceLastRequest < MIN_INTERVAL) {
      await new Promise(resolve => setTimeout(resolve, MIN_INTERVAL - timeSinceLastRequest));
    }

    loading.value = true;
    error.value = null;

    try {
      lastRequestTime = Date.now();

      const response = await fetch(
        `https://nominatim.openstreetmap.org/search?` +
        `format=jsonv2&` +
        `q=${encodeURIComponent(query)}&` +
        `limit=${limit}&` +
        `addressdetails=1&` +
        `accept-language=en`,
        {
          headers: {
            'User-Agent': 'TRAX Solar Energy Platform (development)',
          },
        }
      );

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data: GeocodingResult[] = await response.json();
      results.value = data;
      return data;
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Unknown error occurred';
      error.value = errorMessage;
      console.error('Geocoding error:', err);
      results.value = [];
      return [];
    } finally {
      loading.value = false;
    }
  };

  const getShortDisplayName = (result: GeocodingResult): string => {
    // Extract meaningful parts from display_name
    const parts = result.display_name.split(', ');

    // Take first 2-3 parts for a shorter display
    if (parts.length <= 3) {
      return result.display_name;
    }

    return parts.slice(0, 3).join(', ');
  };

  const clear = () => {
    results.value = [];
    error.value = null;
  };

  return {
    loading,
    error,
    results,
    geocode,
    getShortDisplayName,
    clear,
  };
};
