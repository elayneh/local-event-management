import { ref, watch } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getBookmarkedEvents from "~/graphql/queries/bookmarks/getBookmarkedEvents.gql";
import GetFollowingEvents from "~/graphql/queries/following/getFollowingEvents.gql";
import getTickets from "~/graphql/queries/tickets/getTickets.gql";
import getUserEvents from "~/graphql/queries/events/getUserEvents.gql";

export default function useUserFetchData(user_id) {
  const bookmarkedEvents = ref([]);
  const followingEvents = ref([]);
  const tickets = ref([]);
  const categories = ref([]);
  const tags = ref([]);
  const userEvents = ref([]);

  const pollingInterval = 10000;

  const {
    result: userEventsResult,
    loading: userEventsLoading,
    error: userEventsError,
  } = useQuery(getUserEvents, { id: user_id });

  const {
    result: bookmarkedEventsResult,
    loading: loadingBookmarkedEvents,
    error: bookmarkedEventsError,
  } = useQuery(getBookmarkedEvents, { id: user_id });
  const {
    result: followingEventsResult,
    loading: loadingFollowingEvents,
    error: followingEventsError,
  } = useQuery(
    GetFollowingEvents,
    { id: user_id },
    { pollInterval: pollingInterval }
  );
  const {
    result: ticketsResult,
    loading: loadingTickets,
    error: ticketsError,
  } = useQuery(getTickets, { id: user_id }, { pollInterval: pollingInterval });

  watch(userEventsResult, (newResult) => {
    if (newResult?.userEvents) {
      userEvents.value = newResult.userEvents;
    }
  });

  watch(bookmarkedEventsResult, (newResult) => {
    if (newResult?.bookmarkedEvents) {
      bookmarkedEvents.value = newResult.bookmarkedEvents;
    }
  });

  watch(followingEventsResult, (newResult) => {
    if (newResult?.followingEvents) {
      followingEvents.value = newResult.followingEvents;
    }
  });

  watch(ticketsResult, (newResult) => {
    if (newResult?.tickets) {
      tickets.value = newResult.tickets;
    }
  });

  return {
    followingEvents,
    tickets,
    bookmarkedEvents,
    loadingBookmarkedEvents,
    bookmarkedEventsError,
    userEvents,
    userEventsLoading,
    userEventsError,
    categories,
    tags,
  };
}
