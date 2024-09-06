import { ref } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getEvents from "~/graphql/queries/events/getEvents.gql";
import getCategories from "~/graphql/queries/categories/getCategories.gql";
import getTags from "~/graphql/queries/tags/getTags.gql";

export default function useFetchData(fetchOptions = {}) {
  const defaultOptions = {
    fetchPolicy: "no-cache",
    clientId: "authClient",
    context: {
      headers: {
        "x-hasura-role": "user",
        "x-hasura-is-email-verified": true,
      },
    },
  };

  const options = { ...defaultOptions, ...fetchOptions };

  const events = ref([]);
  const categories = ref([]);
  const tags = ref([]);

  const {
    onResult: onEventResult,
    loading: eventLoading,
    error: eventError,
  } = useQuery(getEvents, {}, options);

  onEventResult(({ data }) => {
    if (data?.events) {
      events.value = data.events;
    }
  });

  const {
    onResult: onCategoryResult,
    loading: loadingCategories,
    error: categoryError,
  } = useQuery(getCategories, {}, options);

  onCategoryResult(({ data }) => {
    if (data?.categories) {
      categories.value = data.categories;
    }
  });

  const {
    onResult: onTagResult,
    loading: loadingTags,
    error: tagError,
  } = useQuery(getTags, {}, options);

  onTagResult(({ data }) => {
    if (data?.tags) {
      tags.value = data.tags;
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
