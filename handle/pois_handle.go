package handle

import (
	"log"
	"pois_debug_client/client"
	"pois_debug_client/config"
	"pois_debug_client/utils"
	"sync"

	"github.com/CESSProject/cess_pois/pois"
	"github.com/pkg/errors"
)

func GenPoisInitialParams() error {
	conf := config.GetConfig()
	minerId, err := utils.ParsingPublickey(conf.MinerId)
	if err != nil {
		return errors.Wrap(err, "generate pois initial params error")
	}
	prover, err := pois.NewProver(conf.K, conf.N, conf.D, []byte(minerId), conf.Space, 32)
	if err != nil {
		return errors.Wrap(err, "generate pois initial params error")
	}
	err = prover.Init(config.GetPoisKey(), pois.Config{})
	//err = prover.Recovery(key, 0, 0, pois.Config{})
	if err != nil {
		return errors.Wrap(err, "generate pois initial params error")
	}

	err = prover.GenerateIdleFileSet()
	if err != nil {
		return errors.Wrap(err, "generate pois initial params error")
	}

	return nil
}

func RequestToGetChallenge(reqNum int) error {
	conf := config.GetConfig()
	minerId, err := utils.ParsingPublickey(conf.MinerId)
	if err != nil {
		return errors.Wrap(err, "request to get commitment challenge error")
	}
	prover, err := pois.NewProver(conf.K, conf.N, conf.D, []byte(minerId), conf.Space, 32)
	if err != nil {
		return errors.Wrap(err, "generate pois initial params error")
	}
	err = prover.Recovery(config.GetPoisKey(), 0, 0, pois.Config{})
	//err = prover.Recovery(key, 0, 0, pois.Config{})
	if err != nil {
		return errors.Wrap(err, "request to get commitment challenge error")
	}

	commits, err := prover.GetIdleFileSetCommits()
	if err != nil {
		return errors.Wrap(err, "request to get commitment challenge error")
	}

	if reqNum <= 0 {
		reqNum = 1
	}
	wg := sync.WaitGroup{}
	wg.Add(reqNum)
	for i := 0; i < reqNum; i++ {
		go func() {
			defer wg.Done()
			chals, err := client.SendCommits(
				minerId, config.GetPoisKey(), 0, 0,
				prover.AccManager.GetSnapshot().Accs.Value, commits,
			)
			if err != nil {
				log.Println("send commits to tee worker error", err)
				return
			}
			log.Println("get challenge success,", len(chals))
		}()
	}
	wg.Wait()
	return nil
}

func RequestToProveCommitProof(num int) error {
	conf := config.GetConfig()
	minerId, err := utils.ParsingPublickey(conf.MinerId)
	if err != nil {
		return errors.Wrap(err, "request to prove commitment proof error")
	}
	prover, err := pois.NewProver(conf.K, conf.N, conf.D, []byte(minerId), conf.Space, 32)
	if err != nil {
		return errors.Wrap(err, "request to prove commitment proof error")
	}
	err = prover.Recovery(config.GetPoisKey(), 0, 0, pois.Config{})
	//err = prover.Recovery(key, 0, 0, pois.Config{})
	if err != nil {
		return errors.Wrap(err, "request to prove commitment proof error")
	}

	commits, err := prover.GetIdleFileSetCommits()
	if err != nil {
		return errors.Wrap(err, "request to prove commitment proof error")
	}
	chals, err := client.SendCommits(
		[]byte(minerId), config.GetPoisKey(), 0, 0,
		prover.AccManager.GetSnapshot().Accs.Value, commits,
	)
	if err != nil {
		return errors.Wrap(err, "request to prove commitment proof error")
	}

	commitProof, accProof, err := prover.ProveCommitAndAcc(chals)
	if err != nil {
		return errors.Wrap(err, "request to prove commitment proof error")
	}

	if num <= 0 {
		num = 1
	} else if num > 1000 {
		num = 1000
	}
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			res, err := client.SendCommitProofs(minerId, commitProof, accProof)
			if err != nil {
				log.Println("request to prove commitment proof error", err)
				return
			}
			log.Println("request to prove commit proof result", res)
		}()
	}
	wg.Wait()
	return nil
}

func RequestToProveSpaceProof(num int) error {
	conf := config.GetConfig()
	minerId, err := utils.ParsingPublickey(conf.MinerId)
	if err != nil {
		return errors.Wrap(err, "request to prove space proof error")
	}
	prover, err := pois.NewProver(conf.K, conf.N, conf.D, []byte(minerId), conf.Space, 32)
	if err != nil {
		return errors.Wrap(err, "request to prove space proof error")
	}
	err = prover.Recovery(config.GetPoisKey(), 0, 256, pois.Config{})
	if err != nil {
		return errors.Wrap(err, "request to prove space proof error")
	}

	verifier := pois.NewVerifier(conf.K, conf.N, conf.D)
	prover.SetChallengeState(
		config.GetPoisKey(),
		prover.GetAccValue(),
		prover.GetFront(),
		prover.GetRear(),
	)
	chal, err := verifier.SpaceChallenges(8)
	if err != nil {
		return errors.Wrap(err, "request to prove space proof error")
	}

	proof, err := prover.ProveSpace(chal, 1, 257)
	if err != nil {
		return errors.Wrap(err, "request to prove space proof error")
	}

	if num <= 0 {
		num = 1
	} else if num > 1000 {
		num = 1000
	}
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			res, err := client.SendSpaceProof(
				minerId, chal, config.GetPoisKey(),
				0, 256, prover.GetAccValue(), proof,
			)
			if err != nil {
				log.Println("request to prove space proof error", err)
				return
			}
			log.Println("request to prove space proof result", res)
		}()
	}
	wg.Wait()
	return nil
}

func RequestToProveDeletionProof(num int) error {
	conf := config.GetConfig()
	minerId, err := utils.ParsingPublickey(conf.MinerId)
	if err != nil {
		return errors.Wrap(err, "request to prove deletion proof error")
	}
	prover, err := pois.NewProver(conf.K, conf.N, conf.D, []byte(minerId), conf.Space, 32)
	if err != nil {
		return errors.Wrap(err, "request to prove deletion proof error")
	}
	err = prover.Recovery(config.GetPoisKey(), 0, 256, pois.Config{})
	if err != nil {
		return errors.Wrap(err, "request to prove deletion proof error")
	}

	proof, err := prover.ProveDeletion(16)
	if err != nil {
		return errors.Wrap(err, "request to prove deletion proof error")
	}

	if num <= 0 {
		num = 1
	} else if num > 1000 {
		num = 1000
	}
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			defer wg.Done()
			res, err := client.SendDeletionProof(
				minerId, config.GetPoisKey(),
				0, 256, prover.GetAccValue(), proof,
			)
			if err != nil {
				log.Println("request to prove deletion proof error", err)
				return
			}
			log.Println("request to prove deletion proof result", res)
		}()
	}
	wg.Wait()
	prover.AccRollback(true)
	return nil
}
