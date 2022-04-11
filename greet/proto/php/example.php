<?php

namespace App\Controller;

use App\Greet\Greeting;
use App\Greet\GreetRequest;
use App\Greet\GreetServiceClient;
use Grpc\ChannelCredentials;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

class TestController extends AbstractController
{
    #[Route('/test', name: 'app_test')]
    public function index(): Response
    {
        $client = new GreetServiceClient(
            'localhost:50051',
            [
                'credentials' => ChannelCredentials::createInsecure(),
            ]
        );

        list($reply, $status) = $client->Greet(
            new GreetRequest(
                [
                    'greeting' => new Greeting(
                        [
                            'title' => 'Jedi',
                            'name' => 'Obiwan'
                        ]
                    )
                ]
            )
        )->wait();

        return $this->json([
            'status' => $status,
            'reply' => $reply->getResult(),
        ]);
    }
}
