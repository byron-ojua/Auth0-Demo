"use client";

import { useEffect, useState } from "react";
import { AppBar, Toolbar, Typography, Button, Avatar } from "@mui/material";
import { useUser } from "@auth0/nextjs-auth0";
import Link from "next/link";

export default function NavBar() {
  const { user } = useUser(); // Get user authentication state
  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);

  // Render nothing until after hydration to avoid mismatch
  if (!mounted) return null;

  return (
    <AppBar position="static">
      <Toolbar>
        {/* App Title */}
        <Typography variant="h6" sx={{ flexGrow: 1 }}>
          Auth0 Demo
        </Typography>

        <Link href="/about" passHref>
          <Button color="inherit">About</Button>
        </Link>

        {/* User Info */}
        {user ? (
          <>
            <Avatar src={user.picture} alt={user.name} sx={{ width: 32, height: 32, mr: 1 }} />
            <Typography variant="body1" sx={{ mr: 2 }}>
              {user.name}
            </Typography>
            <Button color="inherit" href="/auth/logout">
              Logout
            </Button>
          </>
        ) : (
          <Button color="inherit" href="/auth/login">
            Login
          </Button>
        )}
      </Toolbar>
    </AppBar>
  );
}
