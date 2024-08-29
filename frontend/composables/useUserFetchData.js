import { ref, watch } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getBookmarkedEvents from "~/graphql/queries/bookmarks/getBookmarkedEvents.gql";
import GetFollowingEvents from "~/graphql/queries/following/getFollowingEvents.gql";
import getTickets from "~/graphql/queries/tickets/getTickets.gql";
import getUserEvents from "~/graphql/queries/events/getUserEvents.gql";

const bookmarkedEvents = ref([]);
const followingEvents = ref([]);
const tickets = ref([]);
const userEvents = ref([]);

export default function useUserFetchData(user_id) {
  const {
    result: userEventsResult,
    loading: userEventsLoading,
    error: userEventsError,
  } = useQuery(getUserEvents, { id: user_id });

  watch(userEventsResult, (newResult) => {
    if (newResult?.events) {
      userEvents.value = newResult.events;
    }
  });

  const {
    result: bookmarkedEventsResult,
    loading: loadingBookmarkedEvents,
    error: bookmarkedEventsError,
  } = useQuery(getBookmarkedEvents, { id: user_id });

  watch(bookmarkedEventsResult, (newResult) => {
    if (newResult?.events) {
      bookmarkedEvents.value = newResult.events;
    }
  });

  const {
    result: followingEventsResult,
    loading: loadingFollowingEvents,
    error: followingEventsError,
  } = useQuery(GetFollowingEvents, { id: user_id });

  watch(followingEventsResult, (newResult) => {
    console.log("New result: ", newResult);
    if (newResult?.events) {
      followingEvents.value = newResult.events;
    }
  });

  const {
    result: ticketsResult,
    loading: loadingTickets,
    error: ticketsError,
  } = useQuery(getTickets, { id: user_id });

  watch(ticketsResult, (newResult) => {
    if (newResult?.tickets) {
      tickets.value = newResult.tickets;
    }
  });

  return {
    followingEvents,
    tickets,
    bookmarkedEvents,
    userEvents,
  };
}
