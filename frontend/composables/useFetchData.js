import { ref, watch } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getEvents from "~/graphql/queries/events/getEvents.gql";
import getCategories from "~/graphql/queries/categories/getCategories.gql";
import getTags from "~/graphql/queries/tags/getTags.gql";

export default function useFetchData() {
  const events = ref([]);
  const categories = ref([]);
  const tags = ref([]);

  const {
    result: eventResult,
    loading: eventLoading,
    error: eventError,
  } = useQuery(getEvents);
  watch(eventResult, (newResult) => {
    if (newResult?.events) {
      events.value = newResult.events;
    }
  });

  const {
    result: categoryResult,
    loading: loadingCategories,
    error: categoryError,
  } = useQuery(getCategories);
  watch(categoryResult, (newResult) => {
    if (newResult?.categories) {
      categories.value = newResult.categories;
    }
  });

  const {
    result: tagResult,
    loading: loadingTags,
    error: tagError,
  } = useQuery(getTags);
  watch(tagResult, (newResult) => {
    if (newResult?.tags) {
      tags.value = newResult.tags;
    }
  });

  return {
    events,
    categories,
    tags,
    eventLoading,
    loadingCategories,
    loadingTags,
    eventError,
    categoryError,
    tagError,
  };
}
