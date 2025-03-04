"use client";

import { useUser } from '@auth0/nextjs-auth0';
import { useEffect, useState } from 'react';
import { fetchPing } from './requests/ping';
import { buildBaseUrl } from './requests/base';
import { Container } from '@mui/material';

export default function Home() {

  const [message, setMessage] = useState<string | null>(null);

  useEffect(() => {
    console.log("Base URL:", buildBaseUrl());

    fetchPing()
      .then(setMessage)
      .catch((error) => console.error(error));
  }, []);

  return (
    <main>
      <Container>
        <h1>Welcome to the Auth0 Demo</h1>

        {message ? <p>{message}</p> : <p>Loading...</p>}
      </Container>
    </main>
  );
}
