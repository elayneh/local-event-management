<template>
  <div>
    <header>
      <nav>
        <ul>
          <li>
            <NuxtLink to="/">Home</NuxtLink>
          </li>
          <li>
            <NuxtLink to="/about">About</NuxtLink>
          </li>
          <li>
            <NuxtLink to="/users/register">Register</NuxtLink>
          </li>
        </ul>
      </nav>
    </header>
    <slot />
  </div>
</template>
