import { computed } from "vue";
import defaultImage from "~/assets/images/home.png";

export const useEventHelpers = (event) => {
  const truncatedDescription = computed(() => {
    const maxLength = 100;
    return event?.description?.length > maxLength
      ? event.description.slice(0, maxLength) + "..."
      : event?.description || "N/A";
  });

  const eventImages = computed(() => {
    if (event?.event_images) {
      const imagesArray = event.event_images
        .replace(/{|}/g, "")
        .split(",")
        .map((image) => image.trim());
      return imagesArray.length ? imagesArray : [defaultImage];
    }
    return [defaultImage];
  });

  const isFreeText = computed(() => (event?.is_free ? "Free" : "Paid"));

  const freeClass = computed(() =>
    event?.is_free
      ? "bg-green-300 text-white px-2 py-1 rounded"
      : "bg-red-800 text-white px-2 py-1 rounded"
  );

  const formatDate = (dateString) => {
    if (!dateString) return "N/A";
    const date = new Date(dateString);
    return date.toLocaleDateString();
  };

  return {
    truncatedDescription,
    eventImages,
    isFreeText,
    freeClass,
    formatDate,
  };
};
