import { ref } from "vue";

export const useToggleBuy = () => {
  const isBought = ref(false);

  const toggleBuy = () => {
    isBought.value = !isBought.value;
  };

  return { isBought, toggleBuy };
};
