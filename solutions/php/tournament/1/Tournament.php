<?php

/*
 * By adding type hints and enabling strict type checking, code can become
 * easier to read, self-documenting and reduce the number of potential bugs.
 * By default, type declarations are non-strict, which means they will attempt
 * to change the original type to match the type specified by the
 * type-declaration.
 *
 * In other words, if you pass a string to a function requiring a float,
 * it will attempt to convert the string value to a float.
 *
 * To enable strict mode, a single declare directive must be placed at the top
 * of the file.
 * This means that the strictness of typing is configured on a per-file basis.
 * This directive not only affects the type declarations of parameters, but also
 * a function's return type.
 *
 * For more info review the Concept on strict type checking in the PHP track
 * <link>.
 *
 * To disable strict typing, comment out the directive below.
 */

declare(strict_types=1);

class Tournament
{
    private array $teams = [];

    public function tally(string $score): string
    {
        if ($score !== '') {
            foreach (explode("\n", $score) as $line) {
                [$teamA, $teamB, $result] = explode(';', $line);
                $this->teams[$teamA] ??= new TeamScore($teamA);
                $this->teams[$teamB] ??= new TeamScore($teamB);
                switch ($result) {
                    case 'win':
                        $this->teams[$teamA]->addWin();
                        $this->teams[$teamB]->addLoss();
                        break;
                    case 'draw':
                        $this->teams[$teamA]->addDraw();
                        $this->teams[$teamB]->addDraw();
                        break;
                    case 'loss':
                        $this->teams[$teamA]->addLoss();
                        $this->teams[$teamB]->addWin();
                }
            }
        }
        usort($this->teams, static fn(TeamScore $a, TeamScore $b) =>
            [$b->getPoints(), $a->getName()] <=> [$a->getPoints(), $b->getName()]
        );
        return implode("\n", ["Team                           | MP |  W |  D |  L |  P", ...$this->teams]);
    }
}

class TeamScore
{
    private string $team;
    private int    $wins   = 0;
    private int    $losses = 0;
    private int    $draws  = 0;

    public function __construct(string $team)
    {
        $this->team = $team;
    }

    public function addDraw(): void
    {
        $this->draws++;
    }

    public function addLoss(): void
    {
        $this->losses++;
    }

    public function addWin(): void
    {
        $this->wins++;
    }

    public function getMatchesPlayed(): int
    {
        return $this->wins + $this->losses + $this->draws;
    }

    public function getName(): string
    {
        return $this->team;
    }

    public function getPoints(): int
    {
        return 3 * $this->wins + $this->draws;
    }

    public function __toString()
    {
        return sprintf(
            '%-31s|  %d |  %d |  %d |  %d |  %d',
            $this->team,
            $this->getMatchesPlayed(),
            $this->wins,
            $this->draws,
            $this->losses,
            $this->getPoints()
        );
    }
}
