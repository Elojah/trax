import { ref } from 'vue';

export interface ReverseGeocodingResult {
  display_name: string;
  address: {
    house_number?: string;
    road?: string;
    suburb?: string;
    city?: string;
    county?: string;
    state?: string;
    postcode?: string;
    country?: string;
    country_code?: string;
  };
  lat: string;
  lon: string;
  place_id: number;
  osm_type: string;
  osm_id: number;
  type: string;
  importance: number;
}

export const useReverseGeocoding = () => {
  const loading = ref(false);
  const error = ref<string | null>(null);
  const result = ref<ReverseGeocodingResult | null>(null);

  // Rate limiting: Nominatim has a usage policy of max 1 request per second
  let lastRequestTime = 0;
  const MIN_INTERVAL = 1000; // 1 second

  const reverseGeocode = async (lat: number, lng: number): Promise<ReverseGeocodingResult | null> => {
    // Respect rate limiting
    const now = Date.now();
    const timeSinceLastRequest = now - lastRequestTime;
    if (timeSinceLastRequest < MIN_INTERVAL) {
      await new Promise(resolve => setTimeout(resolve, MIN_INTERVAL - timeSinceLastRequest));
    }

    loading.value = true;
    error.value = null;
    result.value = null;

    try {
      lastRequestTime = Date.now();

      const response = await fetch(
        `https://nominatim.openstreetmap.org/reverse?` +
        `format=jsonv2&` +
        `lat=${lat}&` +
        `lon=${lng}&` +
        `zoom=18&` +
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

      const data: ReverseGeocodingResult = await response.json();

      if (data.display_name) {
        result.value = data;
        return data;
      } else {
        throw new Error('No address found for this location');
      }
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Unknown error occurred';
      error.value = errorMessage;
      console.error('Reverse geocoding error:', err);
      return null;
    } finally {
      loading.value = false;
    }
  };

  const formatAddress = (geocodingResult: ReverseGeocodingResult): string => {
    const addr = geocodingResult.address;
    const parts: string[] = [];

    // Build address from components
    if (addr.house_number && addr.road) {
      parts.push(`${addr.house_number} ${addr.road}`);
    } else if (addr.road) {
      parts.push(addr.road);
    }

    if (addr.suburb) {
      parts.push(addr.suburb);
    }

    if (addr.city) {
      parts.push(addr.city);
    } else if (addr.county) {
      parts.push(addr.county);
    }

    if (addr.state) {
      parts.push(addr.state);
    }

    if (addr.postcode) {
      parts.push(addr.postcode);
    }

    if (addr.country) {
      parts.push(addr.country);
    }

    return parts.length > 0 ? parts.join(', ') : geocodingResult.display_name;
  };

  const getShortAddress = (geocodingResult: ReverseGeocodingResult): string => {
    const addr = geocodingResult.address;
    const parts: string[] = [];

    if (addr.road) {
      parts.push(addr.road);
    }

    if (addr.city) {
      parts.push(addr.city);
    } else if (addr.county) {
      parts.push(addr.county);
    }

    if (addr.country) {
      parts.push(addr.country);
    }

    return parts.length > 0 ? parts.join(', ') : geocodingResult.display_name;
  };

  return {
    loading,
    error,
    result,
    reverseGeocode,
    formatAddress,
    getShortAddress,
  };
};
