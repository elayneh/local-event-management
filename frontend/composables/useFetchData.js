import { ref, watch } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getEvents from "~/graphql/queries/events/getEvents.gql";
import getCategories from "~/graphql/queries/categories/getCategories.gql";
import getTags from "~/graphql/queries/tags/getTags.gql";
import getBookmarkedEvents from "~/graphql/queries/bookmarks/getBookmarkedEvents.gql";
import getFollowingEvents from "~/graphql/queries/following/getFollowingEvents.gql";
import getTickets from "~/graphql/queries/tickets/getTickets.gql";

export default function useFetchData(userId) {
  const events = ref([]);
  const categories = ref([]);
  const bookmarkedEvents = ref([]);
  const followingEvents = ref([]);
  const tickets = ref([]);
  const tags = ref([]);

  const pollingInterval = 10000;

  const {
    result: eventResult,
    loading: eventLoading,
    error: eventError,
  } = useQuery(getEvents, null, { pollInterval: pollingInterval });
  watch(eventResult, (newResult) => {
    if (newResult?.events) {
      events.value = newResult.events;
    }
  });

  const {
    result: categoryResult,
    loading: loadingCategories,
    error: categoryError,
  } = useQuery(getCategories, null, { pollInterval: pollingInterval });
  watch(categoryResult, (newResult) => {
    if (newResult?.categories) {
      categories.value = newResult.categories;
    }
  });

  const {
    result: tagResult,
    loading: loadingTags,
    error: tagError,
  } = useQuery(getTags, null, { pollInterval: pollingInterval });
  watch(tagResult, (newResult) => {
    if (newResult?.tags) {
      tags.value = newResult.tags;
    }
  });

  const {
    result: bookmarkedEventsResult,
    loading: loadingBookmarkedEvents,
    error: bookmarkedEventsError,
  } = useQuery(
    getBookmarkedEvents,
    { id: userId },
    { pollInterval: pollingInterval }
  );
  watch(bookmarkedEventsResult, (newResult) => {
    if (newResult?.bookmarkedEvents) {
      bookmarkedEvents.value = newResult.bookmarkedEvents;
    }
  });

  const {
    result: followingEventsResult,
    loading: loadingFollowingEvents,
    error: followingEventsError,
  } = useQuery(
    getFollowingEvents,
    { id: userId },
    { pollInterval: pollingInterval }
  );
  watch(followingEventsResult, (newResult) => {
    if (newResult?.followingEvents) {
      followingEvents.value = newResult.followingEvents;
    }
  });

  const {
    result: ticketsResult,
    loading: loadingTickets,
    error: ticketsError,
  } = useQuery(getTickets, { id: userId }, { pollInterval: pollingInterval });
  watch(ticketsResult, (newResult) => {
    if (newResult?.tickets) {
      tickets.value = newResult.tickets;
    }
  });

  return {
    events,
    categories,
    tags,
    followingEvents,
    tickets,
    bookmarkedEvents,
    eventLoading,
    loadingCategories,
    loadingTags,
    loadingBookmarkedEvents,
    eventError,
    categoryError,
    tagError,
    bookmarkedEventsError,
  };
}
