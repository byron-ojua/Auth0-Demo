import { Container, Typography } from "@mui/material";

export default function AboutPage() {
    return (
        <Container maxWidth="md" sx={{ mt: 4 }}>
            <Typography variant="h3" component="h1" gutterBottom>
                About Us
            </Typography>
        </Container>
    );
}
