import Box from '@mui/material/Box';

export default function Column() {
    return (
        <Box
            height={200}
            width={200}
            my={4}
            display="flex"
            alignItems="center"
            gap={4}
            p={2}
            sx={{ border: '2px solid grey' }}
        >
        </Box>
    )
}