import { ref, computed } from "vue";

export function useTogglePasswordVisibility() {
  const isPasswordVisible = ref(false);

  const togglePasswordVisibility = () => {
    isPasswordVisible.value = !isPasswordVisible.value;
  };

  const passwordFieldType = computed(() =>
    isPasswordVisible.value ? "text" : "password"
  );

  return {
    isPasswordVisible,
    togglePasswordVisibility,
    passwordFieldType,
  };
}
